import axios from 'axios'

const BASE_URL = 'http://localhost:9090/api/v1/auth'
const config = {
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
      "Access-Control-Allow-Headers": "Authorization",
      "Content-Type": "application/json"
    }
};

export const loginUser = async(user) => {
    const requestString = BASE_URL + "/login"
    let response
    try {
      response = await axios.post(requestString, user, config)
    } catch(error) {
      console.log(error)
    }
    return response
}