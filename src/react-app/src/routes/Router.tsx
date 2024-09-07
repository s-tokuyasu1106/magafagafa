import { FC, memo } from "react";
import { Route, Routes } from "react-router-dom";
import {Login} from "../pages/login/Login"
import {Home} from "../pages/home/Home"

export const Router: FC = memo(() => {
    return (
        <Routes>
            <Route path="/" element={<Login />} ></Route>
            <Route path="/home" element={<Home />} ></Route>
        </Routes>
    )
})