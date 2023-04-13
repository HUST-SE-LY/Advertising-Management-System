import React from 'react'
import ReactDOM from 'react-dom/client'
import Home from './pages/home';
import ManagerLogin from './pages/managerLogin';
import CompanyLogin from './pages/companyLogin';
import CompanySignup from './pages/companySignup';
import BackStage from './pages/backStage';
import BackStageHome from './pages/backStage/home';
import BackStageJudge from './pages/backStage/judge';
import BackStageAd from './pages/backStage/ad';
import BackStageCompany from './pages/backStage/company';
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
  },{
    path: '/company-signup',
    element: <CompanySignup></CompanySignup>,
  },{
    path: '/back-stage',
    element: <BackStage></BackStage>,
    children: [
      {
        path: '/back-stage/home',
        element: <BackStageHome></BackStageHome>
      },
      {
        path: '/back-stage/judge',
        element: <BackStageJudge></BackStageJudge>
      },{
        path: '/back-stage/Ads',
        element: <BackStageAd></BackStageAd>
      },{
        path: '/back-stage/companies',
        element: <BackStageCompany></BackStageCompany>
      }
    ]
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
    <RouterProvider router={router} />
)
