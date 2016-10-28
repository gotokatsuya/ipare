# ipare

Image Compare Toolkit  

## Support

### Hash Compare

[x] Average  
[x] Difference

### Pixel Compare

[x] Part
[x] Block

## Usage

### Hash

```go
import (
	"github.com/gotokatsuya/ipare"
	ipareutil "github.com/gotokatsuya/ipare/util"
	"github.com/gotokatsuya/ipare/algorithm/hash"
)

hash := ipare.NewHash()
// Default is DifferenceHash.
// hash.Algorithm = hash.NewAverage()
img1, _ := ipareutil.DecodeImageByPath(imgPath1)
img2, _ := ipareutil.DecodeImageByPath(imgPath2)

const threshold = 5
if distance := hash.Compare(img1, img2); distance <= threshold {
    // similar image
} else {
    // not same image
}
```

### Pixel

```go
import (
	"github.com/gotokatsuya/ipare"
	ipareutil "github.com/gotokatsuya/ipare/util"
	"github.com/gotokatsuya/ipare/algorithm/pixel"
)

pixel := ipare.NewPixel()
// Default is DifferenceHash.
// pixel.Algorithm = pixel.NewBlock()
img1, _ := ipareutil.DecodeImageByPath(imgPath1)
img2, _ := ipareutil.DecodeImageByPath(imgPath2)

const threshold = 10
if distance := pixel.Compare(img1, img2); distance <= threshold {
    // similar image
} else {
    // not same image
}
```


## Test

```bash
$ go test github.com/gotokatsuya/ipare -v
```

