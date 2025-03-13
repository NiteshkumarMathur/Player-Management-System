package main

import (
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/rlimit"
	"github.com/vishvananda/netlink"
)

const (
	ifaceName   = "eth0"
	defaultPort = 4040
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <port>")
		os.Exit(1)
	}

	port := uint16(defaultPort)
	fmt.Sscanf(os.Args[1], "%d", &port)

	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		fmt.Printf("failed to remove memory limit: %v\n", err)
		os.Exit(1)
	}

	// Load the eBPF program.
	spec, err := ebpf.LoadCollectionSpec("drop_tcp_port.o")
	if err != nil {
		fmt.Printf("failed to load eBPF program: %v\n", err)
		os.Exit(1)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		fmt.Printf("failed to create eBPF collection: %v\n", err)
		os.Exit(1)
	}
	defer coll.Close()

	// Update the port map with the desired port.
	portMap := coll.Maps["port_map"]
	key := uint32(0)
	if err := portMap.Update(&key, &port, ebpf.UpdateAny); err != nil {
		fmt.Printf("failed to update port map: %v\n", err)
		os.Exit(1)
	}

	// Attach the eBPF program to the interface.
	iface, err := netlink.LinkByName(ifaceName)
	if err != nil {
		fmt.Printf("failed to get interface %s: %v\n", ifaceName, err)
		os.Exit(1)
	}

	link := netlink.LinkAttrs{Index: iface.Attrs().Index}
	prog := coll.Programs["drop_tcp_port"]

	xdp := netlink.LinkXdp{
		Program: prog,
		Flags:   netlink.XDP_FLAGS_UPDATE_IF_NOEXIST,
	}
	if err := netlink.LinkSetXdp(iface, &xdp); err != nil {
		fmt.Printf("failed to attach xdp program: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully attached eBPF program to interface %s\n", ifaceName)
}
