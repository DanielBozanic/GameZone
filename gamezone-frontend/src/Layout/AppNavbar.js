import "../Assets/css/app-navbar.css";
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
                <Link to="/">Video Games</Link>
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink>
                <Link to="/">Consoles</Link>
              </NavLink>
            </NavItem>
            <UncontrolledDropdown inNavbar nav>
              <DropdownToggle caret nav>
                Hardware components
              </DropdownToggle>
              <DropdownMenu right>
                <DropdownItem>
                  <Link to="/">Graphics cards</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">RAMs</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Motherboards</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Processors</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Solid state drives</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Hard disk drives</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Monitors</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Power supply units</Link>
                </DropdownItem>
              </DropdownMenu>
            </UncontrolledDropdown>
            <UncontrolledDropdown inNavbar nav>
              <DropdownToggle caret nav>
                Accessories
              </DropdownToggle>
              <DropdownMenu right>
                <DropdownItem>
                  <Link to="/">Headphones</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Keyboards</Link>
                </DropdownItem>
                <DropdownItem>
                  <Link to="/">Mice</Link>
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
