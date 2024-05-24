import { useContext, useEffect, useState } from "react"
import { GoInfo } from "react-icons/go";
import { IoMailSharp } from "react-icons/io5";
import { FaMapLocationDot } from "react-icons/fa6";
import { MdOutlinePassword } from "react-icons/md";
import { FaUser } from "react-icons/fa";
import { CgMenuMotion } from "react-icons/cg";


const RegisterComponent = () => {
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [password1, setPassword1] = useState('')
    const [password2, setPassword2] = useState('')

    const submit = (e) => {
        e.preventDefault()
        console.log(name, email, password1, password2)
    }

    return (
        <section className="">
            <div className="container mx-auto">
                <form className="text-center mx-auto w-1/6 mb-0 mt-32">
                    <label htmlFor="name" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <FaUser className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setName(e.target.value)} type="text" name="name" required id="name" placeholder="John" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>

                    <label htmlFor="email" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <IoMailSharp className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setEmail(e.target.value)} type="email" name="email" id="email" required placeholder="example@gmail.com" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>

                    <label htmlFor="password" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <MdOutlinePassword className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setPassword1(e.target.value)} type="password" name="password" required id="password" placeholder="Password" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>

                    <label htmlFor="repeatPassword" className="relative text-gray-400 focus-within:text-gray-600 block mb-2">
                        <MdOutlinePassword className="pointer-events-none w-5 h-5 translate-x-1/4 -translate-y-1/2 absolute top-1/2 left-1"/>
                        <input onChange={(e) => setPassword2(e.target.value)} type="password" name="repeatPassword" required id="repeatPassword" placeholder="Repeat password" className="form-input w-full pl-10 py-2 rounded-sm"/>
                    </label>
                    
                    <button onClick={submit} className="w-full rounded-sm bg-slate-600 hover:bg-slate-500 active:bg-slate-700 text-white px-5 py-2 uppercase shadow-2xl">
                        Register
                    </button>
                </form>

                <div className="w-1/6 mx-auto text-center py-2 flex items-center">
                    <hr className="w-1/2 border-1/2 border-slate-600"/>
                    <span className="px-2 mx-auto">or</span>
                    <hr className="w-1/2 border-1/2 border-slate-600"/>
                </div>

                <div className="w-1/6 mx-auto text-center">
                    <p className=""><a href="/login" className="text-blue-500 underline underline-offset-2">Sign in</a>, if already have an account.</p>
                </div>
            </div>
        </section>
    )
}

export default RegisterComponent