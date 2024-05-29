package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {

	// DATA PENGUKURAN
	// Massa 1, Pegas 1 Pada Udara
	deltaX1 := []float64{5, 5, 4.5, 4.3, 4, 4, 3.8, 3.4, 3, 3}
	time1 := []float64{0.7, 1.5, 2.3, 3.2, 3.9, 4.7, 5.6, 6.6, 7.4, 8.3}

	// Massa 2, Pegas 1 Pada Udara
	deltaX2 := []float64{5.8, 5.5, 5.5, 5.2, 5.1, 5, 4.7, 4.3, 4, 4}
	time2 := []float64{0.8, 1.5, 2.2, 3, 3.8, 4.7, 5.4, 6.2, 7.1, 8}

	// Massa 1, Pegas 1 Pada Air
	deltaX3 := []float64{5.3, 5.3, 5.1, 5, 5, 4.7, 4.5, 4.4, 4, 4}
	time3 := []float64{0.7, 1.3, 2.1, 2.8, 3.6, 4.4, 5.2, 6, 6.7, 7.5}

	// Massa 2, Pegas 1 Pada Air
	deltaX4 := []float64{5.7, 5.2, 5, 4.7, 4.5, 4.5, 4.1, 3.6, 3.4, 3}
	time4 := []float64{0.8, 1.7, 2.4, 3.1, 3.9, 4.7, 5.6, 6.3, 7.1, 8}

	// Massa 1, Pegas 2 Pada Udara
	deltaX5 := []float64{5.1, 4.8, 4.6, 4.6, 4, 3.9, 3.6, 3.4, 3.1, 3}
	time5 := []float64{0.7, 1.4, 2.2, 3, 3.9, 4.7, 5.5, 6.3, 7.2, 8}

	// Massa 2, Pegas 2 Pada Udara
	deltaX6 := []float64{4, 3.5, 3.2, 3.3, 2.7, 2.7, 2.5, 2.1, 2, 1.8}
	time6 := []float64{0.8, 1.8, 2.8, 3.8, 4.7, 5.7, 6.7, 7.5, 8.4, 9.4}

	// Massa 1, Pegas 2 Pada Air
	deltaX7 := []float64{4.4, 4.4, 4, 3.8, 3.7, 3.5, 3.3, 3.1, 3, 3}
	time7 := []float64{0.9, 1.7, 2.4, 3.2, 4, 4.8, 5.6, 6.3, 7.2, 8}

	// Massa 2, Pegas 2 Pada Air
	deltaX8 := []float64{4, 4, 3.5, 3.2, 3.2, 3, 2.8, 2.7, 2.5, 2.3}
	time8 := []float64{1, 1.9, 3, 3.9, 4.9, 5.9, 7, 8, 9, 9.9}

	k1 := 6.64  // konstanta pegas 1
	k2 := 5.97  // konstanta pegas 2
	m1 := 0.164 // massa 1
	m2 := 0.231 // massa 2
	r  := 0.025  // jari-jari beban

	perhitungan()
	plotGrafik1(deltaX1, time1, "Variasi Massa 1 Pegas 1 Pada Udara", "regresi_eksponensial_1.png")
	plotGrafik1(deltaX2, time2, "Variasi Massa 2 Pegas 1 Pada Udara", "regresi_eksponensial_2.png")
	plotGrafik1(deltaX3, time3, "Variasi Massa 1 Pegas 1 Pada Air", "regresi_eksponensial_3.png")
	plotGrafik1(deltaX4, time4, "Variasi Massa 2 Pegas 1 Pada Air", "regresi_eksponensial_4.png")
	plotGrafik1(deltaX5, time5, "Variasi Massa 1 Pegas 2 Pada Udara", "regresi_eksponensial_5.png")
	plotGrafik1(deltaX6, time6, "Variasi Massa 2 Pegas 2 Pada Udara", "regresi_eksponensial_6.png")
	plotGrafik1(deltaX7, time7, "Variasi Massa 1 Pegas 2 Pada Air ", "regresi_eksponensial_7.png")
	plotGrafik1(deltaX8, time8, "Variasi Massa 2 Pegas 2 Pada Air", "regresi_eksponensial_8.png")

	b1 := menghitungB(deltaX1, time1, m1) // 0.5534851269355424
	b2 := menghitungB(deltaX2, time2, m2) // 0.8380877735222986
	b3 := menghitungB(deltaX3, time3, m1) // 0.5668796121769734
	b4 := menghitungB(deltaX4, time4, m2) // 0.8402521541823444
	b5 := menghitungB(deltaX5, time5, m1) // 0.5558247334910178
	b6 := menghitungB(deltaX6, time6, m2) // 0.6712974079157135
	b7 := menghitungB(deltaX7, time7, m1) // 0.5063764427932589
	b8 := menghitungB(deltaX8, time8, m2) // 0.6714470785136224

	omega1 := menghitungOmega(k1, m1, b1) // omega 1: 40.452625
	omega2 := menghitungOmega(k2, m2, b2) // omega 2: 25.780412
	omega3 := menghitungOmega(k1, m1, b3) // omega 3: 40.450901
	omega4 := menghitungOmega(k2, m1, b4) // omega 4: 36.312188
	omega5 := menghitungOmega(k1, m2, b5) // omega 5: 28.719401
	omega6 := menghitungOmega(k2, m1, b6) // omega 6: 36.344860
	omega7 := menghitungOmega(k1, m1, b7) // omega 7: 40.458360
	omega8 := menghitungOmega(k2, m2, b8) // omega 8: 25.803259

	// Grafik
	plotGrafik2(deltaX1, time1, omega1, "massa 1", "Pegas 1", "Udara", "grafik_osilasi_1.png")
	plotGrafik2(deltaX2, time2, omega2, "massa 2", "Pegas 1", "Udara", "grafik_osilasi_2.png")
	plotGrafik2(deltaX3, time3, omega3, "massa 1", "Pegas 1", "Air", "grafik_osilasi_3.png")
	plotGrafik2(deltaX4, time4, omega4, "massa 2", "Pegas 1", "Air", "grafik_osilasi_4.png")
	plotGrafik2(deltaX5, time5, omega5, "massa 1", "Pegas 2", "Udara", "grafik_osilasi_5.png")
	plotGrafik2(deltaX6, time6, omega6, "massa 2", "Pegas 2", "Udara", "grafik_osilasi_6.png")
	plotGrafik2(deltaX7, time7, omega7, "massa 1", "Pegas 2", "Air", "grafik_osilasi_7.png")
	plotGrafik2(deltaX8, time8, omega8, "massa 2", "Pegas 2", "Air", "grafik_osilasi_8.png")

	hitungViskositas(deltaX1, time1, r, m1) // Massa 1, Pegas 1 udara
	hitungViskositas(deltaX2, time2, r, m2) // Massa 2, pegas 1 udara
	hitungViskositas(deltaX3, time3, r, m1) // Massa 1, pegas 1 air
	hitungViskositas(deltaX4, time4, r, m2) // Massa 2, pegas 1 air
	hitungViskositas(deltaX5, time5, r, m1) // Massa 1, pegas 2 udara
	hitungViskositas(deltaX6, time6, r, m2) // Massa 2, pegas 2 udara
	hitungViskositas(deltaX7, time7, r, m1) // Massa 1, pegas 2 air
	hitungViskositas(deltaX8, time8, r, m2) // massa 2, pegas 2 air

}

func perhitungan() {
	// Variabel diketahui
	m := 61e-3 // massa (kg)
	g := 9.8   // percepatan gravitasi (m/s^2)
	x1 := 9e-2 // jarak (m)
	x2 := 10e-2

	k1 := (m * g) / x1
	k2 := (m * g) / x2

	// Menampilkan nilai k1 & k2
	fmt.Printf("Nilai k1 = %.3f N/m\n", k1) // Nilai k1 = 6.642 N/m
	fmt.Printf("Nilai k2 = %.3f N/m\n", k2) // Nilai k2 = 5.978 N/m

}

func plotGrafik1(deltaX, time []float64, judul, fileName string) {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Grafik Amplitudo Osilasi Terhadap Waktu Variasi %s ", judul)
	p.Title.TextStyle.Font.Size = vg.Points(20)
	p.X.Label.Text = "Waktu (s)"
	p.X.Label.TextStyle.Font.Size = vg.Points(18)
	p.Y.Label.Text = "Amplitudo (cm)"
	p.Y.Label.TextStyle.Font.Size = vg.Points(18)
	p.Legend.Top = true

	// Membuat titik-titik sebagai data asli
	points := make(plotter.XYs, len(time))
	for i := range time {
		points[i].X = time[i]
		points[i].Y = deltaX[i]
	}

	// Error handling
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Points(3)
	scatter.GlyphStyle.Color = plotutil.Color(2)
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}

	// Estimasi koefisien regresi eksponensial
	a, b := nilaiRegresi(deltaX, time)

	// Data regresi eksponensial
	line := plotter.NewFunction(func(x float64) float64 { return a * math.Exp(b*x) })
	line.Color = plotutil.Color(1)
	line.Width = vg.Points(2)

	// Menambahkan data dan regresi eksponensial ke plot
	p.Add(scatter, line)
	p.Legend.Add("Data Asli", scatter)
	p.Legend.Add(fmt.Sprintf("y = %.3fe^(%.3fx)", a, b), line)
	p.Add(plotter.NewGrid())
	p.Legend.TextStyle.Font.Size = vg.Points(18)

	// Menyimpan plot ke file dengan bentuk .png
	if err := p.Save(10*vg.Inch, 8*vg.Inch, fileName); err != nil {
		log.Panic(err)
	}
}

// nilaiRegresi mengestimasi parameter a dan b dari data deltaX dan time
func nilaiRegresi(deltaX, time []float64) (float64, float64) {
	n := float64(len(deltaX))
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0

	for i := range deltaX {
		logY := math.Log(deltaX[i])
		sumX += time[i]
		sumY += logY
		sumXY += time[i] * logY
		sumXX += time[i] * time[i]
	}

	b := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	a := (sumY - b*sumX) / n
	a = math.Exp(a)

	return a, b
}

// menghitungOmega menghitung nilai omega
func menghitungOmega(k, m, b float64) float64 {
	return math.Sqrt(math.Pow(k/m, 2) - math.Pow(b/(2*m), 2))
}

// Fungsi untuk mencari koefisien b
func menghitungB(deltaX, time []float64, m float64) float64 {
	lndeltaX := make([]float64, len(deltaX))
	for i := range deltaX {
		lndeltaX[i] = math.Log(deltaX[i])
	}
	slope, _ := stat.LinearRegression(time, lndeltaX, nil, false)
	b := 2 * m * slope
	return b
}

func plotGrafik2(deltaX, time []float64, omega float64, massa, pegas, medium, fileName string) {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Grafik Osilasi Terhadap Waktu Variasi %s %s pada %s", massa, pegas, medium)
	p.Title.TextStyle.Font.Size = vg.Points(18)
	p.X.Label.Text = "Waktu (s)"
	p.X.Label.TextStyle.Font.Size = vg.Points(16)
	p.Y.Label.Text = "Amplitudo (cm)"
	p.Y.Label.TextStyle.Font.Size = vg.Points(16)
	p.Legend.Top = true

	// Data asli
	points := make(plotter.XYs, len(time))
	for i := range time {
		points[i].X = time[i]
		points[i].Y = deltaX[i]
	}
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Points(3)
	scatter.GlyphStyle.Color = plotutil.Color(5)
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}

	a, b := nilaiRegresi(deltaX, time)

	line := plotter.NewFunction(func(x float64) float64 { return a * math.Exp(b*x) })
	line.Color = plotutil.Color(1)
	line.Width = vg.Points(2)

	// Oscillation plot
	oscillation := make(plotter.XYs, 1000)
	for i := range oscillation {
		t := float64(i) * time[len(time)-1] / 1000
		oscillation[i].X = t
		oscillation[i].Y = a * math.Exp(b*t) * math.Cos(omega*t)
	}
	lineOscillation, err := plotter.NewLine(oscillation)
	if err != nil {
		log.Panic(err)
	}
	lineOscillation.Color = plotutil.Color(2)

	// Menambahkan komponen ke plot
	p.Add(lineOscillation)
	p.Add(scatter, line)
	p.Legend.Add("Data Asli", scatter)
	p.Legend.Add(fmt.Sprintf("y = %.3fe^(%.3fx)", a, b), line)
	p.Legend.Add("Osilasi", lineOscillation)
	p.Legend.TextStyle.Font.Size = vg.Points(14)

	// Menyimpan plot
	if err := p.Save(8*vg.Inch, 6*vg.Inch, fileName); err != nil {
		log.Panic(err)
	}
}

func hitungViskositas(deltaX, time []float64, r, m float64) {
	pi := math.Pi
	b := menghitungB(deltaX, time, m)
	viscosity := b * 0.001 / (6 * pi * r)
	fmt.Printf("Viskositas: %.5f Pa.S\n", viscosity)
}