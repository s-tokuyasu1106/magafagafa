import {FC, memo} from "react"
import { LogoutButton } from "../common/logout/LogoutButton";

export const Home: FC = memo(()=>{
    return (
        <>
            <p>Homeページ</p>
            <LogoutButton></LogoutButton>
        </>
    );
})