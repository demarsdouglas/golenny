package lenny

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

var TestLenny = Lenny{
    leftEye:  "O",
    rightEye: "O",
    leftEnd:  "(",
    rightEnd: ")",
    mouth:    "w",
}

func TestRender(t *testing.T) {
    Convey("Given a Lenny struct", t, func() {
        l := TestLenny
        Convey("When the render function is called", func() {
            r := l.Render()
            Convey("The value of the render should equal the values in the struct", func() {
                So(r, ShouldEqual, "(OwO)")
            })
        })
    })
}
