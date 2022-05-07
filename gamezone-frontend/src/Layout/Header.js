import logo from "../Assets/images/logo.PNG";
import "../Assets/css/header.css";
import { Row, Col, Container } from "reactstrap";

const Header = () => {
  return (
    <div className="header">
      <Container>
        <Row>
          <Col>
            <img src={logo} alt="Logo" className="responsive-img-header" />
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Header;
