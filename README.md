httpjson
==
unmarsh post data from http
-
##go server
***
struct define whith json tag
***
    type test struct {

        Iint int                `json:"int"`
        Sstring string      `json:"string`
        Slice []string       `json:"slice"`
    }
##html
    <form method="POST" action="assign_profile">
        <table>
        <tr><td>int:</td><td><input type="text" name="int" value="5"/></td></tr>
        <tr><td>string:</td><td><input type="text" name="string" value="abcd"/></td></tr>
        <tr><td>slice:</td><td><input type="text" name="slice" value="ab cd ef"/></td></tr>
        </table>
    </form>

##main
***
slice only support string now
***
    func Parse(r *http.Request){
        t := &test{}
        json.unmarshal(r, t)
        fmt.Println(*t)
        //t.Iint    5
        //t.Sstring abcd
        //t.Slice ["ab","cd", "ef"]
    }
