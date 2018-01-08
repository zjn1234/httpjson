# httpjson

## unmarshal post data from http
***
if you html form has a lot of items, then the library fit you, it can simplify code struct
***
### case

>##### struct defination
    type test struct {
        Iint int            `json:"int"`
        Sstring string      `json:"string`
        Slice []string      `json:"slice"`
    }
>#### html defination
    <form method="POST" action="assign_profile">
        <table>
        <tr><td>int:</td><td><input type="text" name="int" value="5"/></td></tr>
        <tr><td>string:</td><td><input type="text" name="string" value="abcd"/></td></tr>
        <tr><td>slice:</td><td><input type="text" name="slice" value="ab cd ef"/></td></tr>
        </table>
    </form>
>#### then you can
	func supplyService (w ResponseWriter, r *http.Request){
    	var t &test{}
        err := httpjson.Unmarshal(r, t)
        ........
    }
    
#### rules
1. slice only support int float32、float64、string、bool、type
2. it look up json tag first as key, if not found, replace with lower-case field name, then find value from httpRequest using key