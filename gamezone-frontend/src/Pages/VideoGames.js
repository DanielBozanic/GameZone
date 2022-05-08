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
import cn from "classnames";
import { useState, useEffect } from "react";

const VideoGames = () => {
  const [videoGames, setVideoGames] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [pageCount, setPageCount] = useState([]);
  const pageSize = 8;

  const handleClick = (e, index) => {
    e.preventDefault();
    setCurrentPage(index);
  };

  useEffect(() => {
    const getVideoGames = () => {
      axios
        .get(
          `http://localhost:7000/api/products/videoGames?page=${currentPage}&pageSize=${pageSize}`
        )
        .then((res) => {
          setVideoGames(res.data.VideoGames);
          setPageCount(Math.ceil(res.data.PageCount / pageSize));
        })
        .catch((err) => {});
    };

    getVideoGames();
  }, [currentPage]);

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
        <Row
          className={cn(
            "pagination",
            videoGames.length < 4
              ? "pagination-padding"
              : "pagination-padding-normal"
          )}
        >
          <Col md="12">
            <Pagination size="lg">
              <PaginationItem disabled={currentPage <= 1}>
                <PaginationLink
                  onClick={(e) => handleClick(e, currentPage - 1)}
                  previous
                />
              </PaginationItem>

              {[...Array(pageCount)].map((page, i) => (
                <PaginationItem active={i === currentPage - 1} key={i}>
                  <PaginationLink onClick={(e) => handleClick(e, i + 1)}>
                    {i + 1}
                  </PaginationLink>
                </PaginationItem>
              ))}
              <PaginationItem disabled={currentPage - 1 >= pageCount - 1}>
                <PaginationLink
                  onClick={(e) => handleClick(e, currentPage + 1)}
                  next
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
