package factory

// ChainableModifier creates a factory that produces an initial value from initFactory, then lets you change modifiers to modify that value.
// You're then returned the value once all modifiers have ran over it. This can be particularly useful when building mock data in tests.
//
//	func example() {
//		chain := ChainableModifier(func() *pb.MyRequest {
//			return &pb.MyRequest{Foo: true, Bar: false}
//		})
//
//		firstRequest := chain() // returns &pb.MyRequest{Foo: true, Bar: false}
//
//		secondRequestWithBarTrue := chain(func(v *pb.MyRequest) {
//			v.Bar = true
//		}) // returns &pb.MyRequest{Foo: true, Bar: true}
//
//		thirdRequestWithValuesSwapped := chain(func(v *pb.MyRequest) {
//			v.Foo = false
//		}, func(v *pb.MyRequest) {
//			v.Bar = true
//		}) // returns &pb.MyRequest{Foo: false, Bar: true}
//	}
func ChainableModifier[T any](initFactory func() T) func(modifiers ...func(T)) T {
	return func(modifiers ...func(T)) T {
		out := initFactory()

		for _, fn := range modifiers {
			fn(out)
		}

		return out
	}
}
