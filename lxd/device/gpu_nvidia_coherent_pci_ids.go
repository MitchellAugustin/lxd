package device

import (
	"strings"
)

// isNvidiaCoherentGPU returns true if the PCI vendor/product ID pair is a known NVIDIA coherent-memory GPU.
func isNvidiaCoherentGPU(vendorID string, productID string) bool {
	if !strings.EqualFold(vendorID, "10de") {
		return false
	}

	// Only Grace superchip IDs qualify; discrete GPUs using the same die do not.
	switch strings.ToLower(productID) {
	case "2342", "2348", "2941": // GH200 120GB/480GB, GH200 144GB HBM3e, GB200.
		return true
	default:
		return false
	}
}
