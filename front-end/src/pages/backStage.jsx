import { NavLink, Outlet } from "react-router-dom"

function BackStage() {
  return (<div className="w-screen h-screen grid grid-cols-[1fr_5fr] p-[2rem]">
    <div className="w-3/4 max-w-[8rem] h-full bg-blue-400 rounded-3xl shadow-xl shadow-blue-300/30 flex flex-col justify-center items-center gap-[5rem] p-3"> 
      <NavLink to={'/back-stage/home'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':''}>
        <img src="/src/assets/home.svg" alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>
      
      <NavLink to={'/back-stage/judge'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':''}>
        <img src="/src/assets/judge.svg" alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>

      <NavLink to={'/back-stage/companies'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':''}>
        <img src="/src/assets/company.svg" alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>

      <NavLink to={'/back-stage/ads'} className={({isActive, isPending}) => isActive?'w-fit bg-blue-600 rounded-2xl shadow-lg shadow-blue-700/60 transition-all':isPending?'':''}>
        <img src="/src/assets/ad.svg" alt="" className="w-1/3 block mx-[auto] my-[17%]" />
      </NavLink>
    </div>
    <Outlet></Outlet>
  </div>)
}

export default BackStage