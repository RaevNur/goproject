go run . "hello world" --color=red
go run . "1 + 1 = 2" --color=green
go run . "(%&) ??" --color=yellow
go run . "This is Letters" "--color=RGB(200,100,100) in={2-15}"
go run . "This is Letters" "--color=RGB(100,200,100) in={2}"
go run . "This is Letters" "--color=RGB(100,100,200) substr={e}"
go run . "HeY GuYs" "--color=orange substr={GuYs}"
go run . "RGB()" --color=blue
go run . "RGBrgbRgBRGb" --color=cyan
go run . "RGBrgbRgBRGb 123 {}" "--color=cyan symbols={1{}"
go run . "RGBrgbRgBRGb {}" "--color=cyan symbols={{}}"
go run . "asdfaghWDASDC" "--color=RGB(100,250,200) symbols={aS}"
go run . "asdfaghWD1AS2D3C }|}" "--color=RGB(100,250,200) symbols={1a2S3|}"
