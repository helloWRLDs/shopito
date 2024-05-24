import LoginComponent from './component/LoginComponent'
import {BrowserRouter, Routes, Route} from 'react-router-dom'
import './index.css'
import RegisterComponent from './component/RegisterComponent'
import MainComponent from './component/MainComponent'
import { HeaderComponent, FooterComponent } from './component/StaticComponents'

function App() {
  return (
    <>
      <div className="app">
        <BrowserRouter>
          <HeaderComponent />
            <Routes>
              <Route path="/" element={<MainComponent />}/>
              <Route path="/login" element= {<LoginComponent />}/>
              <Route path="/register" element= {<RegisterComponent />}/>
            </Routes>
        </BrowserRouter>
      </div>
    </>
  )
}

export default App
