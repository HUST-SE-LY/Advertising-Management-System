import { NavLink, Outlet } from "react-router-dom"
import HomeSvg from '/src/assets/home.svg'
import JudgeSvg from '/src/assets/judge.svg'
import CompanySvg from '/src/assets/company.svg'
import AdSvg from '/src/assets/ad.svg'

function BackStage() {
  return (<div className="w-screen h-screen grid grid-cols-[1fr_6fr] p-[2rem]">
    <div className="w-3/4 max-w-[8rem] ml-[2rem] h-full bg-blue-400 rounded-3xl shadow-xl shadow-blue-300/30 flex flex-col justify-center items-center gap-[5rem] p-3"> 
      <NavLink to={'/back-stage/home'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':'hover:bg-blue-500 rounded-2xl transition-all'}>
        <img src={HomeSvg} alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>
      
      <NavLink to={'/back-stage/judge'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all ':isPending?'':'hover:bg-blue-500 rounded-2xl transition-all'}>
        <div className="relative before:w-[10px] before:h-[10px] before:rounded-full before:bg-red-500 before:absolute before:z-40 before:top-0 before:right-1/3 before:animate-ping">
          <img src={JudgeSvg} alt="" className="w-1/3 block mx-[auto] my-[17%] " />
        </div>
        
      </NavLink>

      <NavLink to={'/back-stage/companies'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':'hover:bg-blue-500 rounded-2xl transition-all'}>
        <img src={CompanySvg} alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>

      <NavLink to={'/back-stage/ads'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':'hover:bg-blue-500 rounded-2xl transition-all'}>
        <img src={AdSvg} alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>
    </div>
    <Outlet></Outlet>
  </div>)
}

export default BackStage