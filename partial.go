package httprouter
import "net/http"

type PartialRouter struct {
	trees  map[string]*node
}

type StdHandle func(http.ResponseWriter, *http.Request)

// Make sure the PartialRouter conforms with the http.Handler interface
var _ http.Handler = NewPartial()

// New returns a new initialized PartialRouter.
// Path auto-correction, including trailing slashes, is enabled by default.
func NewPartial() *PartialRouter {
	return &PartialRouter{ }
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (r *PartialRouter) GET(path string, handle StdHandle) {
	r.Handle("GET", path, handle)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func (r *PartialRouter) HEAD(path string, handle StdHandle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func (r *PartialRouter) OPTIONS(path string, handle StdHandle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (r *PartialRouter) POST(path string, handle StdHandle) {
	r.Handle("POST", path, handle)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (r *PartialRouter) PUT(path string, handle StdHandle) {
	r.Handle("PUT", path, handle)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (r *PartialRouter) PATCH(path string, handle StdHandle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (r *PartialRouter) DELETE(path string, handle StdHandle) {
	r.Handle("DELETE", path, handle)
}

// Handle registers a new request handle with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *PartialRouter) Handle(method, path string, handle StdHandle) {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}


	//TODO: Partial matches should NOT be allowed to have parameters or wildcards

	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}

	root.addRoute(path, func(w http.ResponseWriter, r *http.Request, ps Params) {
		handle(w, r)
	})
}

// ServeHTTP makes the router implement the http.Handler interface.
func (r *PartialRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if root := r.trees[req.Method]; root != nil {
		if handle, ps, _ := root.getValue(path); handle != nil {
			handle(w, req, ps)
			return
		}
	}
}
