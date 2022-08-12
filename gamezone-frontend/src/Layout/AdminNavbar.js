import {
	Nav,
	NavbarToggler,
	UncontrolledDropdown,
	DropdownToggle,
	DropdownMenu,
	DropdownItem,
	Collapse,
	NavItem,
	NavLink,
	Container,
	Col,
} from "reactstrap";
import { useState } from "react";
import { Link } from "react-router-dom";
import * as authService from "../Auth/AuthService";

const AdminNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	const signOut = () => {
		authService.removeToken();
	};

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Container fluid style={{ padding: "0" }}>
					<Nav className="me-auto" navbar>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/profile">Profile</Link>
								</NavLink>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<UncontrolledDropdown inNavbar nav>
									<DropdownToggle caret nav>
										User management
									</DropdownToggle>
									<DropdownMenu right>
										<DropdownItem>
											<Link
												className="drop-down-link"
												to="/createAdminAndEmployee"
											>
												Add employee/admin
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/registeredUsers">
												View registered users
											</Link>
										</DropdownItem>
									</DropdownMenu>
								</UncontrolledDropdown>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/businessReports">Business reports</Link>
								</NavLink>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/viewNews">News</Link>
								</NavLink>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<NavLink href="/signIn" onClick={signOut}>
									Sign Out
								</NavLink>
							</NavItem>
						</Col>
					</Nav>
				</Container>
			</Collapse>
		</>
	);
};

export default AdminNavbar;
