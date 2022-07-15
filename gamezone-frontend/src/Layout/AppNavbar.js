import "../Assets/css/navbar.css";
import { Link } from "react-router-dom";
import {
	Navbar,
	Nav,
	NavItem,
	UncontrolledDropdown,
	DropdownMenu,
	DropdownItem,
	DropdownToggle,
	NavbarToggler,
	Collapse,
	NavLink,
} from "reactstrap";
import { useState } from "react";

const AppNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	return (
		<div className="navbar">
			<Navbar color="dark" expand="md" dark>
				<NavbarToggler onClick={toggleNavbar} className="mr-2" />
				<Collapse isOpen={!collapsed} navbar>
					<Nav className="me-auto" navbar>
						<NavItem>
							<NavLink>
								<Link to="/videoGames">Video Games</Link>
							</NavLink>
						</NavItem>
						<NavItem>
							<NavLink>
								<Link to="/consoles">Consoles</Link>
							</NavLink>
						</NavItem>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								Hardware components
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/graphicsCards">
										Graphics cards
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/rams">
										RAMs
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/motherboards">
										Motherboards
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/processors">
										Processors
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/ssds">
										Solid state drives
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/hdds">
										Hard disk drives
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/monitors">
										Monitors
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/psus">
										Power supply units
									</Link>
								</DropdownItem>
							</DropdownMenu>
						</UncontrolledDropdown>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								Accessories
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/headphones">
										Headphones
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/keyboards">
										Keyboards
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/mice">
										Mice
									</Link>
								</DropdownItem>
							</DropdownMenu>
						</UncontrolledDropdown>
					</Nav>
				</Collapse>
			</Navbar>
		</div>
	);
};

export default AppNavbar;
