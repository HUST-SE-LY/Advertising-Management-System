import HomeSvg from '/src/assets/home.svg'
import CompanySvg from '/src/assets/company.svg'
import { NavLink, Outlet } from 'react-router-dom'

function User() {
  return <div>
    <TabBar></TabBar>
    <Outlet></Outlet>
  </div>
}

function TabBar() {
  return <div className="w-full px-[2rem] py-[1rem] flex items-center rounded-[0_0_2rem_2rem] gap-[2rem] bg-blue-300 shadow-lg shadow-blue-600/10 ">
    <p className="text-xl font-bold text-gray-700 tracking-widest">欢迎回来!</p>
    <div className="ml-[auto] flex gap-[1rem]">
      <NavLink to="/user/home" className={({isActive}) => isActive?'bg-blue-600 shadow-xl shadow-blue-600/20 transition-all rounded-full p-[1rem]':'hover:bg-blue-400 transition-all rounded-full p-[1rem]'}>
        <img src={HomeSvg} className="h-[2rem]" alt="" />
      </NavLink>
      <NavLink to="/user/detail" className={({isActive}) => isActive?'bg-blue-600 shadow-xl shadow-blue-600/20 transition-all rounded-full p-[1rem]':'hover:bg-blue-400 transition-all rounded-full p-[1rem]'}>
        <img src={CompanySvg} className="h-[2rem]" alt="" />
      </NavLink>
    </div>
  </div>
}

export default User