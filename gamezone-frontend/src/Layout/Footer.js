import { Row, Col, Container } from "reactstrap";
import logo from "../Assets/images/logo.PNG";

const Footer = () => {
  return (
    <div className="footer">
      <Container>
        <Row>
          <Col>
            <img src={logo} alt="Logo" className="responsive-img-footer" />
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Footer;
