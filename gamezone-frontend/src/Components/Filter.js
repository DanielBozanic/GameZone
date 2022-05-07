import "../Assets/css/filter.css";
import {
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Form,
  FormGroup,
  Label,
  Input,
  Row,
  Col,
  Container,
} from "reactstrap";
import { useState } from "react";

const Filter = () => {
  const [modal, setModal] = useState(false);
  const toggle = () => setModal(!modal);
  return (
    <>
      <Container>
        <Row className="main-container">
          <Col md="2">
            <Button className="filter-button" onClick={toggle}>
              Filter
            </Button>
          </Col>
          <Col md="10">
            <Row>
              <Col md="10" style={{ textAlign: "right" }}>
                <Input
                  className="search-bar"
                  type="text"
                  placeholder="Search"
                />
              </Col>
              <Col md="2">
                <Button className="search-btn">Search</Button>
              </Col>
            </Row>
          </Col>
        </Row>
      </Container>
      <Modal isOpen={modal} toggle={toggle}>
        <ModalHeader toggle={toggle}>Filter</ModalHeader>
        <ModalBody>
          <Form>
            <FormGroup>
              <Label>Filter 1</Label>
              <Input
                id="exampleSelectMulti"
                multiple
                name="selectMulti"
                type="select"
              >
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
              </Input>
            </FormGroup>
            <FormGroup>
              <Label>Filter 2</Label>
              <Input
                id="exampleSelectMulti2"
                multiple
                name="selectMulti"
                type="select"
              >
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
              </Input>
            </FormGroup>
            <FormGroup>
              <Label>Filter 3</Label>
              <Input
                id="exampleSelectMulti3"
                multiple
                name="selectMulti"
                type="select"
              >
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
              </Input>
            </FormGroup>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button className="confirm-filter-btn">Filter</Button>
        </ModalFooter>
      </Modal>
    </>
  );
};

export default Filter;
