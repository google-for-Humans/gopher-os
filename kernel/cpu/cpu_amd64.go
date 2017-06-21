package cpu

// FlushTLBEntry flushes a TLB entry for a particular virtual address.
func FlushTLBEntry(virtAddr uintptr)

// SwitchPDT sets the root page table directory to point to the specified
// physical address and flushes the TLB.
func SwitchPDT(pdtPhysAddr uintptr)

// ActivePDT returns the physical address of the currently active page table.
func ActivePDT() uintptr
