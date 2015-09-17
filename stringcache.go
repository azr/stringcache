//Pkg that implements string caching to reduce mem use
//
// We count on the fact that golang strings are immutable.
//
// Also remember that golang strings are simply language structs that point to bytes.
// Meaning that if you try to compare the address of cached strings, they will differ.
package stringcache

type Getter interface {
	//Get(foo) will store `foo` on first call
	//and will return its own `foo` on future calls
	Get(string) string
}
