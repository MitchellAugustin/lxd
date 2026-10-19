package device

import (
	"testing"
)

func TestIsNvidiaCoherentGPU(t *testing.T) {
	tests := []struct {
		name      string
		vendorID  string
		productID string
		want      bool
	}{
		{name: "known Grace Hopper (GH200 120GB/480GB)", vendorID: "10de", productID: "2342", want: true},
		{name: "known Grace Hopper (GH200 144G HBM3e)", vendorID: "10de", productID: "2348", want: true},
		{name: "known Grace Blackwell (HGX GB200)", vendorID: "10de", productID: "2941", want: true},
		{name: "known vendor, plain (non-Grace) Hopper board", vendorID: "10de", productID: "2335", want: false},
		{name: "known vendor unknown product", vendorID: "10de", productID: "0000", want: false},
		{name: "unknown vendor", vendorID: "1002", productID: "2342", want: false},
		{name: "empty vendor", vendorID: "", productID: "2342", want: false},
		{name: "empty product", vendorID: "10de", productID: "", want: false},
		{name: "both empty", vendorID: "", productID: "", want: false},
		{name: "case insensitive match", vendorID: "10DE", productID: "2342", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNvidiaCoherentGPU(tt.vendorID, tt.productID)
			if got != tt.want {
				t.Errorf("isNvidiaCoherentGPU(%q, %q) = %v, want %v", tt.vendorID, tt.productID, got, tt.want)
			}
		})
	}
}
