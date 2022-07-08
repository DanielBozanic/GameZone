import logo from "../Assets/images/logo.PNG";
import "../Assets/css/header.css";
import "../Assets/css/app-navbar.css";
import { Link } from "react-router-dom";
import { Navbar, NavbarBrand } from "reactstrap";
import UnauthenticatedNavbar from "./UnauthenticatedNavbar";
import UserNavbar from "./UserNavbar";
import * as authService from "../Auth/AuthService";

const Header = () => {
	const navbar = () => {
		if (authService.getRole() === "ROLE_USER") {
			return <UserNavbar />;
		} else {
			return <UnauthenticatedNavbar />;
		}
	};

	return (
		<Navbar className="header" color="dark" expand="md" dark>
			<NavbarBrand>
				<Link to="/">
					<img src={logo} alt="Logo" className="responsive-img-header" />
				</Link>
			</NavbarBrand>
			{navbar()}
		</Navbar>
	);
};

export default Header;
