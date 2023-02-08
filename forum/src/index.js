// import React from 'react';
// import ReactDOM from 'react-dom/client';
// import './index.css';

// import LoginPage from './routes/LoginPage';
// import SignUpPage from './routes/Signup';
// import HomePage from './routes/HomePage'

// import { BrowserRouter as Router, Routes, Route, useLoaderData } from "react-router-dom";
// import {
//   createBrowserRouter,
//   RouterProvider,
  
// } from "react-router-dom";





// const router = createBrowserRouter([
//   {
//     path: "/",
//     element: <LoginPage/>,
//   },
//   {
//     path:"/signup",
//     element: <SignUpPage/>,
//   },
//   {
//     path:"/home",
//     element: <HomePage />
//   }
// ]);

import React from 'react';
import { BrowserRouter, Route, Routes} from "react-router-dom";
import ReactDOM from 'react-dom';

import LoginPage from './routes/LoginPage';
import SignUpPage from './routes/Signup';
import HomePage from './routes/HomePage';

const App = () => (
  <BrowserRouter>
    <Routes>
      <Route exact path="/" element={<LoginPage/>} />
      <Route path="/signup" element={<SignUpPage/>} />
      <Route path="/home" element={<HomePage/>} />
    </Routes>
  </BrowserRouter>
);

export default App;
const container = document.getElementById('root');
//const root = ReactDOM.createRoot(document.getElementById('root'));
ReactDOM.render(<App />, container);

// root.render(
//   <React.StrictMode>
//     <App />
//     {/* <RouterProvider router={router} /> */}
//   </React.StrictMode>
// );