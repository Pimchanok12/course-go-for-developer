
/*
--- example Maps.java ---
public class Maps {
	public static void main(String[] args) {
		Map<String, String> map = new HashMap<>();
		map.put("id", "1");
		map.put("airlineCode", "FD");
	}}

--- example Maps.go ---

func DemoMap() {
	// var m map[string]string
	// m = make(map[string]string)

	// m := make(map[string]string)
	m := map[string]string{
		"id":         "1",
		"airlineCode": "FD",
	}

	m["id"] = "1"
	m["airlineCode"] = "FD"

	fmt.Println(m["id"])
}


*/
