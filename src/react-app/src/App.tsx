import React from 'react';
// import logo from './logo.svg';
import './App.css';
import { ChakraProvider } from "@chakra-ui/react";
import {BrowserRouter} from "react-router-dom"
import {Router} from "./routes/Router"
// import { useAuth0 } from "@auth0/auth0-react";
// import LoginButton from './pages/login/LoginButton';



function App() {
  return (
    <ChakraProvider>
      <BrowserRouter>
      {/* <div> */}
        {/* <header> */}
          {/* <LoginButton></LoginButton> */}
          {/* <LogoutButton></LogoutButton> */}
        {/* </header> */}
      {/* </div> */}
      <Router />
      </BrowserRouter>
    </ChakraProvider>
  );
}

export default App;
