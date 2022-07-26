import { Nav, NavbarToggler, Collapse, NavItem, NavLink } from "reactstrap";
import { useState } from "react";
import { Link } from "react-router-dom";
import * as authService from "../Auth/AuthService";

const UserNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	const signOut = () => {
		authService.removeToken();
	};

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Nav className="me-auto" navbar>
					<NavItem>
						<NavLink>
							<Link to="/shoppingCart">Shopping Cart</Link>
						</NavLink>
					</NavItem>
					<NavItem>
						<NavLink>
							<Link to="/viewNews">News</Link>
						</NavLink>
					</NavItem>
					<NavItem>
						<NavLink href="/signIn" onClick={signOut}>
							Sign Out
						</NavLink>
					</NavItem>
				</Nav>
			</Collapse>
		</>
	);
};

export default UserNavbar;
