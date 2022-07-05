import logo from "../Assets/images/logo.PNG";
import "../Assets/css/header.css";
import { Navbar, NavbarBrand } from "reactstrap";

const Header = () => {
	return (
		<Navbar className="header" color="dark" expand="md" dark>
			<NavbarBrand>
				<img src={logo} alt="Logo" className="responsive-img-header" />
			</NavbarBrand>
		</Navbar>
	);
};

export default Header;
