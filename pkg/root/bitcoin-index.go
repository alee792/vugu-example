package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

var _ vugu.ComponentType = (*BitcoinIndex)(nil)

func (comp *BitcoinIndex) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*BitcoinIndexData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "bitcoin-index"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		if data.isLoading {
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Loading...", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		if len(data.priceIndex.PriceIndex) > 0 {
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				{
					parent := n
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Updated: ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					n.InnerHTML = fmt.Sprint(data.priceIndex.Time.Updated)
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "ul", DataAtom: vugu.VGAtom(42754), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				{
					parent := n
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					for key, value := range data.priceIndex.PriceIndex {
						_, _ = key, value
						n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "li", DataAtom: vugu.VGAtom(45570), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
						{
							parent := n
							n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
							n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
							n.InnerHTML = fmt.Sprint(key)
							n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
							n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
							n.InnerHTML = fmt.Sprint(fmt.Sprint(value.Symbol, value.RateFloat))
							n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
						}
					}
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		// @click = { data.HandleClick(event) }
		{
			var i_ interface{} = data
			idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
			var i2_ interface{} = data.HandleClick
			i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
			n.SetDOMEventHandler("click", vugu.DOMEventHandler{
				ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
				Method:                reflect.ValueOf(data).MethodByName("HandleClick"),
				Args:                  []interface{}{event},
			})
		}
		if false {
			// force compiler to check arguments for type safety
			data.HandleClick(event)
		}
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Fetch Bitcoin Price Index", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "ul", DataAtom: vugu.VGAtom(42754), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			for i := 0; i < 10; i++ {
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "my-line", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "file-name", Val: "example.txt"}}}
				parent.AppendChild(n)
				n.Props = vugu.Props{
					"line-number": i,
				}
				{
					parent := n
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    \n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
				}
			}
		}
	}
	return
}

type BitcoinIndex struct {}

func (ct *BitcoinIndex) NewData(props vugu.Props) (interface{}, error) { return &BitcoinIndexData{}, nil }

func init() { vugu.RegisterComponentType("bitcoin-index", &BitcoinIndex{}) }
