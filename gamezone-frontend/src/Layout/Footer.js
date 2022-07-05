import logo from "../Assets/images/logo.PNG";
import "../Assets/css/footer.css";
import { Navbar, NavbarBrand } from "reactstrap";

const Footer = () => {
	return (
		<Navbar className="footer" color="dark" expand="md" dark>
			<NavbarBrand>
				<img src={logo} alt="Logo" className="responsive-img-footer" />
			</NavbarBrand>
		</Navbar>
	);
};

export default Footer;
