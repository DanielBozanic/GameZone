import * as productAPI from "../APIs/ProductMicroservice/product_api";
import AppNavbar from "../Layout/AppNavbar";
import Search from "../Components/Search";
import ProductsView from "../Components/ProductsView";
import { Row, Col, Container, Card, CardHeader, CardTitle } from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Navigation } from "swiper";
import { Helmet } from "react-helmet";
import * as helperFunctions from "../Utils/HelperFunctions";
import * as authService from "../Auth/AuthService";

const Main = () => {
	const [products, setProducts] = useState([]);
	const [mainPageProducts, setMainPageProducts] = useState([]);
	const [popularProducts, setPopularProducts] = useState([]);
	const [recommendedProducts, setRecommendedProducts] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const [searchTerm, setSearchTerm] = useState("");
	const pageSize = 12;

	const handleClick = (e, index) => {
		e.preventDefault();
		setCurrentPage(index);
	};

	useEffect(() => {
		if (searchTerm !== "") {
			getProductsSearch();
			getPageCountSearch();
		} else {
			setProducts([]);
			getMainPageProducts();
			getPopularProducts();
			if (authService.isUser()) {
				getRecommendedProducts();
			}
		}
	}, [currentPage, searchTerm]);

	const getProductsSearch = () => {
		axios
			.get(
				`${productAPI.SEARCH_BY_NAME}
					?page=${currentPage}&pageSize=${pageSize}&name=${searchTerm}`
			)
			.then((res) => {
				setProducts(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCountSearch = () => {
		axios
			.get(`${productAPI.GET_NUMBER_OF_RECORDS_SEARCH}?name=${searchTerm}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getMainPageProducts = () => {
		axios
			.get(`${productAPI.GET_MAIN_PAGE_PRODUCTS}`)
			.then((res) => {
				setMainPageProducts(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getPopularProducts = () => {
		axios
			.get(`${productAPI.GET_POPULAR_PRODUCTS}`)
			.then((res) => {
				setPopularProducts(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getRecommendedProducts = () => {
		axios
			.get(`${productAPI.GET_RECOMMENDED_PRODUCTS}`)
			.then((res) => {
				setRecommendedProducts(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const viewProductDetail = (product) => {
		return helperFunctions.getProductDetailRoute(product);
	};

	const onSearchClick = (st) => {
		setSearchTerm(st);
		setCurrentPage(1);
	};

	return (
		<>
			<Helmet>
				<title>Home | GameZone</title>
			</Helmet>
			<Container>
				<Row style={{ display: "flex" }}>
					<Search onSearchClick={onSearchClick} searchPlaceholder={"Search"} />
				</Row>
			</Container>
			<AppNavbar />
			{searchTerm === "" && (
				<Container>
					{recommendedProducts.length > 0 && (
						<Row className="swiper-row">
							<Col>
								<Card style={{ margin: "0" }}>
									<CardHeader>
										<CardTitle className="title" tag="h4">
											Because you like {recommendedProducts[0].Manufacturer}{" "}
											products
										</CardTitle>
									</CardHeader>
									<Swiper slidesPerView={3}>
										{recommendedProducts.map((product) => {
											return (
												<SwiperSlide>
													<a
														className="main-page-product-link"
														href={viewProductDetail(product)}
													>
														<img
															className="main-page-product-image"
															src={product.Image.Content}
														/>
													</a>
												</SwiperSlide>
											);
										})}
									</Swiper>
								</Card>
							</Col>
						</Row>
					)}
					{mainPageProducts.length > 0 && (
						<Row className="swiper-row">
							<Col>
								<Card style={{ margin: "0" }}>
									<CardHeader>
										<CardTitle className="title" tag="h4">
											Featuring
										</CardTitle>
									</CardHeader>
									<Swiper
										slidesPerView={3}
										autoplay={{
											delay: 3500,
											disableOnInteraction: false,
										}}
										navigation={true}
										modules={[Autoplay, Navigation]}
									>
										{mainPageProducts.map((product) => {
											return (
												<SwiperSlide>
													<a
														className="main-page-product-link"
														href={viewProductDetail(product)}
													>
														<img
															className="main-page-product-image"
															src={product.Image.Content}
														/>
													</a>
												</SwiperSlide>
											);
										})}
									</Swiper>
								</Card>
							</Col>
						</Row>
					)}
					{popularProducts.length > 0 && (
						<Row className="swiper-row">
							<Col>
								<Card style={{ margin: "0" }}>
									<CardHeader>
										<CardTitle className="title" tag="h4">
											Popular
										</CardTitle>
									</CardHeader>
									<Swiper slidesPerView={3}>
										{popularProducts.map((product) => {
											return (
												<SwiperSlide>
													<a
														className="main-page-product-link"
														href={viewProductDetail(product)}
													>
														<img
															className="main-page-product-image"
															src={product.Image.Content}
														/>
													</a>
												</SwiperSlide>
											);
										})}
									</Swiper>
								</Card>
							</Col>
						</Row>
					)}
				</Container>
			)}
			<ProductsView
				products={products}
				currentPage={currentPage}
				pageCount={pageCount}
				handleClick={handleClick}
			/>
		</>
	);
};

export default Main;
