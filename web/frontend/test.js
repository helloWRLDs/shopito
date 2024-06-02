import axios from "axios";

const config = {
    headers: {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
      }
}

const main = async() => {
    const response = await axios.get("http://localhost:9090/api/v1/products/1", config)
    console.log(response.data)
}

main()