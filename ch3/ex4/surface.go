// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
    "fmt"
    "io"
    "log"
    "math"
    "net/http"
    "os"
)

const (
    width, height = 600, 320             // canvas size in pixels
    cells         = 100                  // number of grid cells
    xyrange       = 30.0                 // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange  // pixels per x or y unit
    zscale        = height * 0.4         // pixels per z unit
    angle         = math.Pi / 6          // angle of x, y axes (=30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30 deg), cos(30 deg)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "web" {
        handler := func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "image/svg+xml")
            genSVG(w)
        }
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe("localhost:8000", nil))
        return
    } 

    if len(os.Args) > 1 && os.Args[1] == "file" {
        gf, _ := os.Create("out.svg")
        genSVG(gf)
        gf.Close()
        return
    }
    genSVG(os.Stdout)
}

func genSVG(out io.Writer) {
    output := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
                "style='stroke: grey; fill: white; stroke-width: 0.7' "+
                "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            var ok bool
            var ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz float64
            ax, ay, az, ok = corner(i+1, j)
            if ok == false {
                continue
            }
            bx, by, bz, ok = corner(i, j)
            if ok == false {
                continue
            }
            cx, cy, cz, ok = corner(i, j+1)
            if ok == false {
                continue
            }
            dx, dy, dz, ok = corner(i+1, j+1)
            if ok == false {
                continue
            }

            z := (az + bz + cz + dz) / 4.0
            red := uint8(0xFF)
            green := uint8(0xFF)
            blue := uint8(0xFF)
            switch {
            case z > float64(0.0):
                green = uint8(0xFF - (0xFF * z))
                blue = uint8(0xFF - (0xFF * z))
            case z < float64(0.0):
                red = uint8(0xFF - (0xFF * -z))
                green = uint8(0xFF - (0xFF * -z))
            }
            output += fmt.Sprintf("<polygon style='fill: #%02X%02X%02X' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
                red, green, blue, ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    output += fmt.Sprintln("</svg>")
    fmt.Fprintf(out, output)
}

func corner(i, j int) (float64, float64, float64, bool) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z, ok := f(x, y)
    if ok == false {
        return 0, 0, 0, false
    }

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z, true
}

func f(x, y float64) (float64, bool) {
    r := math.Hypot(x,y) // distance from (0,0)
    val := math.Sin(r) / r
    if math.IsNaN(val) == true {
        return 0, false
    }
    return val, true
}

func g(x, y float64) (float64, bool) {
    z := math.Hypot(x,y)
    if math.IsNaN(z) == true {
        return 0, false
    }
    return z, true
}
