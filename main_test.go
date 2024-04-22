package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		name            string
		hargaTotal      float64
		metodePembayaran string
		dicicil         bool
		wantErr         bool
	}{
		{
			name:            "Case 1: Harga nol",
			hargaTotal:      0,
			metodePembayaran: "transfer",
			dicicil:         false,
			wantErr:         true,
		},
		{
			name:            "Case 2: Metode pembayaran tidak dikenali",
			hargaTotal:      100000,
			metodePembayaran: "gopay",
			dicicil:         false,
			wantErr:         true,
		},
		{
			name:            "Case 3: Cicilan tidak memenuhi syarat",
			hargaTotal:      400000,
			metodePembayaran: "credit",
			dicicil:         true,
			wantErr:         true,
		},
		{
			name:            "Case 4: Credit harus dicicil",
			hargaTotal:      600000,
			metodePembayaran: "credit",
			dicicil:         false,
			wantErr:         true,
		},
		{
			name:            "Case 5: Pembayaran berhasil",
			hargaTotal:      600000,
			metodePembayaran: "credit",
			dicicil:         true,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.hargaTotal, tt.metodePembayaran, tt.dicicil)
			if (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
