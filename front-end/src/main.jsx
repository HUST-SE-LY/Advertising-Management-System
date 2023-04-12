import React from 'react'
import ReactDOM from 'react-dom/client'
import Home from './pages/home';
import ManagerLogin from './pages/managerLogin';
import CompanyLogin from './pages/companyLogin';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import './index.css'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home></Home>,
  },
  {
    path: '/manager-login',
    element: <ManagerLogin></ManagerLogin>
  },{
    path: '/company-login',
    element: <CompanyLogin></CompanyLogin>
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
    <RouterProvider router={router} />
)
