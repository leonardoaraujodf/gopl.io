// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
    "fmt"
    "math"
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
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            var ok bool
            var ax, ay, bx, by, cx, cy, dx, dy float64
            ax, ay, ok = corner(i+1, j)
            if ok == false {
                continue
            }
            bx, by, ok = corner(i, j)
            if ok == false {
                continue
            }
            cx, cy, ok = corner(i, j+1)
            if ok == false {
                continue
            }
            dx, dy, ok = corner(i+1, j+1)
            if ok == false {
                continue
            }
            fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z, ok := f(x, y)
    if ok == false {
        return 0, 0, false
    }

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, true
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
