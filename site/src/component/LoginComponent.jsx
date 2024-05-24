import React, { useContext, useEffect, useState } from "react"
import { GoInfo } from "react-icons/go";
import { IoMailSharp } from "react-icons/io5";
import { MdOutlinePassword } from "react-icons/md";
import { loginUser } from "../services/authService";
import { getProduct } from "../services/prodcuService";

const LoginComponent = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')

    const submit = async(e) => {
        e.preventDefault()
        const user = {
            email: email,
            password: password
        }
        const response = await getProduct("1")
        console.log(response.data)
    }

    const validateEmail = () => {
        let valid = true
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
        if (email.trim()) {
            setError('')
            if (emailRegex.test(email)) {
                setError('')
            } else {
                setError('Wrong email format'); 
                valid=false;
            }
        } else {
            setError('Email is required'); 
            valid=false;
        }
        return valid
    }

    const validatePassword = () => {
        // let valid = validateEmail()
        // const passwordRegex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/
        // if (password.trim()) {
        //     setError('')
        //     if (passwordRegex.test(password)) {setError('')}
        //     else {setError('Min 8 symbols(1 letter, 1 num)'); valid=false}
        // }
        // else setError("Password is required"); valid=false;
        // return valid
    }

    return (
        <section className="">
            <div className="container mx-auto">
                <form className="text-center mx-auto w-1/6 mb-0 mt-52">
                    <label htmlFor="email" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <IoMailSharp className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setEmail(e.target.value)} onBlur={validateEmail} type="email" name="email" id="email" placeholder="example@gmail.com" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>

                    <label htmlFor="password" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <MdOutlinePassword className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setPassword(e.target.value)} onBlur={validatePassword} type="password" name="password" id="password" placeholder="Password" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>
                    {error &&
                        <div className="pointer-events-none w-full flex items-center px-2 py-2 rounded-sm bg-red-600 text-white mb-2"><GoInfo className="mr-3 w-5 h-5"/>{error}</div>
                    }

                    <button onClick={submit} className="w-full rounded-sm bg-slate-600 hover:bg-slate-500 active:bg-slate-700 text-white px-5 py-2 uppercase shadow-2xl">
                        Login
                    </button>
                </form>

                <div className="w-1/6 mx-auto text-center py-2 flex items-center">
                    <hr className="w-1/2 border-1/2 border-slate-600"/>
                    <span className="px-2 mx-auto">or</span>
                    <hr className="w-1/2 border-1/2 border-slate-600"/>
                </div>

                <div className="w-1/6 mx-auto text-center">
                    <p className=""><a href="/signup" className="text-blue-500 underline underline-offset-2">Sign up</a>, if you haven't yet.</p>
                </div>
            </div>
        </section>
    )
}

export default LoginComponent