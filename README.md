# improved-couscous
golang with testing


    $ go test ./...
    $ go test ./... -v

**test with httptest**

    func TestGet(t *testing.T){
        r := httptest.NewRequest(http.MethodGet, "/", nil)
        w := httptest.NewRecorder()

        ServerHTTP(w, r)
        
        if w.Body.String() == "" {
            t.Errorf("")
        }

        if w.Code == http.StatusOK {
            t.Errorf("")
        }
    }

    func TestGet(t *testing.T){
        r := httptest.NewRequest(http.MethodGet, "/", nil)
        w := httptest.NewRecorder()
    }


**run test**

    $ go test ./...

result:

    ok      github.com/ilmsg/improved-couscous/cat  (cached)
    ok      github.com/ilmsg/improved-couscous/hello        (cached)
    ok      github.com/ilmsg/improved-couscous/note 0.006s

**run test with description**

    $ go test -v ./...

result:

    === RUN   TestServeHTTP
    --- PASS: TestServeHTTP (0.00s)
    PASS
    ok      github.com/ilmsg/improved-couscous/cat  (cached)
    === RUN   TestHelloHandler
    --- PASS: TestHelloHandler (0.00s)
    PASS
    ok      github.com/ilmsg/improved-couscous/hello        (cached)
    === RUN   TestList_Hi
    --- PASS: TestList_Hi (0.00s)
    === RUN   TestList
    --- PASS: TestList (0.00s)
    PASS
    ok      github.com/ilmsg/improved-couscous/note 0.006s
