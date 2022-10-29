// Package try
//
// For the public APIs, returns error is a good choice.
//
// However, in the internal calls, err != nil checks are everywhere,
// even though we don't need to know about what exactly it is,
// and just want to propagate it to caller.
//
// The try package provide abilities to propagate error(via panic)
// in internal calls and convert them back to error in exposed
// functions.
//
//	func ExposedFunc() string {
//	    return try.Catching(func() string {
//	        if !satisfied {
//	            try.Throw("illegal argument")
//	        }
//	    })
//	}
package try
