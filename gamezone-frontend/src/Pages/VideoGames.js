import "../Assets/css/video-games.css";
import AppNavbar from "../Layout/AppNavbar";
import Filter from "../Components/Filter";
import {
  CardText,
  CardImg,
  CardTitle,
  CardBody,
  Card,
  Row,
  Col,
  Container,
  Pagination,
  PaginationItem,
  PaginationLink,
} from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";

const VideoGames = () => {
  const [videoGames, setVideoGames] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [pageCount, setPageCount] = useState([]);
  const [pageSize, setPageSize] = useState(8);

  useEffect(() => {
    getPageCount();
    getVideoGames();
  }, [currentPage]);

  const handleClick = (e, index) => {
    e.preventDefault();
    setCurrentPage(index);
  };

  const getPageCount = () => {
    axios
      .get(
        `http://localhost:7000/api/products/videoGames/getPageCount?pageSize=${pageSize}`
      )
      .then((res) => {
        let pc = Math.ceil(res.data / pageSize);
        setPageCount(pc);
      })
      .catch((err) => {});
  };

  const getVideoGames = () => {
    axios
      .get(
        `http://localhost:7000/api/products/videoGames?page=${currentPage}&pageSize=${pageSize}`
      )
      .then((res) => {
        setVideoGames(res.data);
      })
      .catch((err) => {});
  };

  return (
    <>
      <Filter />
      <AppNavbar />
      <Container>
        <Row className="card-row">
          {videoGames.map((product, index) =>
            index < 4 ? (
              <Col style={{ paddingTop: "5px" }} md={3}>
                <Card className="card-with-image">
                  <CardImg
                    className="card-image"
                    alt="Card image cap"
                    src={product.Product.Image}
                  />
                  <CardBody>
                    <CardTitle tag="h5">{product.Product.Name}</CardTitle>
                    <CardText>{product.Product.Price}RSD</CardText>
                  </CardBody>
                </Card>
              </Col>
            ) : (
              ""
            )
          )}
        </Row>
        <Row className="card-row">
          {videoGames.map((product, index) =>
            index > 3 ? (
              <Col style={{ paddingTop: "5px" }} md={3}>
                <Card className="card-with-image">
                  <CardImg
                    className="card-image"
                    alt="Card image cap"
                    src={product.Product.Image}
                  />
                  <CardBody>
                    <CardTitle tag="h5">{product.Product.Name}</CardTitle>
                    <CardText>{product.Product.Price}RSD</CardText>
                  </CardBody>
                </Card>
              </Col>
            ) : (
              ""
            )
          )}
        </Row>
        <Row className="pagination">
          <Col md="12">
            <Pagination size="lg">
              <PaginationItem disabled={currentPage <= 1}>
                <PaginationLink
                  onClick={(e) => handleClick(e, currentPage - 1)}
                  previous
                  href="#"
                />
              </PaginationItem>

              {[...Array(pageCount)].map((page, i) => (
                <PaginationItem active={i === currentPage - 1} key={i}>
                  <PaginationLink
                    onClick={(e) => handleClick(e, i + 1)}
                    href="#"
                  >
                    {i + 1}
                  </PaginationLink>
                </PaginationItem>
              ))}
              <PaginationItem disabled={currentPage - 1 >= pageCount - 1}>
                <PaginationLink
                  onClick={(e) => handleClick(e, currentPage + 1)}
                  next
                  href="#"
                />
              </PaginationItem>
            </Pagination>
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default VideoGames;
