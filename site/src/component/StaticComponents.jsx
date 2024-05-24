
export const HeaderComponent = () => {
    return (
        <header className="bg-slate-600 text-white py-4">
        <div className="container mx-auto flex items-center justify-between">
            <div>
                <a href="/" className="text-xl font-bold">Shopito</a>
            </div>
            <nav>
                <ul className="flex space-x-4">
                    <li><a href="/" className="hover:text-gray-300">Home</a></li>
                    <li><a href="#" className="hover:text-gray-300">Shop</a></li>
                    <li><a href="#" className="hover:text-gray-300">About</a></li>
                    <li><a href="#" className="hover:text-gray-300">Contact</a></li>
                </ul>
            </nav>
            <div className="flex items-center space-x-4">
                <a href="/login" className="hover:text-gray-300">Login</a>
                <a href="/register" className="hover:text-gray-300">Sign Up</a>
                <a href="#" className="hover:text-gray-300">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 15l7-7 7 7"></path>
                    </svg>
                </a>
            </div>
        </div>
    </header>
    )
}

export const FooterComponent = () => {
    return (
        <h1>Footer</h1>
    )
}