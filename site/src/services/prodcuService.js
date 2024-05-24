import axios from 'axios'

const BASE_URL = 'http://localhost:9090/api/v1/products'
const config = {
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
      "Access-Control-Allow-Headers": "Authorization",
      "Content-Type": "application/json"
    }
};

export const getProduct = async(id) => {
    const response = await axios.get(BASE_URL + `/${id}`, config)
    console.log(response.data)
}