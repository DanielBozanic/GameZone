import * as productAPI from "../APIs/ProductMicroservice/product_api";
import AppNavbar from "../Layout/AppNavbar";
import Search from "../Components/Search";
import ProductsView from "../Components/ProductsView";
import { Row, Col, Container } from "reactstrap";
import axios from "axios";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Pagination, Navigation } from "swiper";
import ProductsMainModal from "../Components/ProductsMainModal";
import * as authService from "../Auth/AuthService";
import * as helperFunctions from "../Utils/HelperFunctions";

const Main = () => {
	const [products, setProducts] = useState([]);
	const [mainPageProducts, setMainPageProducts] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const [searchTerm, setSearchTerm] = useState("");
	const pageSize = 8;

	const navigate = useNavigate();

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

	const viewProductDetail = (product) => {
		navigate(helperFunctions.getProductDetailRoute(product));
	};

	const onSearchClick = (st) => {
		setSearchTerm(st);
		setCurrentPage(1);
	};

	return (
		<>
			<Container>
				<Row style={{ display: "flex" }}>
					{authService.isEmployee() && (
						<Col md="2">
							<ProductsMainModal />
						</Col>
					)}
					<Search onSearchClick={onSearchClick} />
				</Row>
			</Container>
			<AppNavbar />
			{searchTerm === "" && mainPageProducts.length > 0 && (
				<Container>
					<Row className="swiper-row">
						<Col>
							<Swiper
								slidesPerView={3}
								spaceBetween={10}
								autoplay={{
									delay: 3500,
									disableOnInteraction: false,
								}}
								pagination={{
									clickable: true,
								}}
								navigation={true}
								modules={[Autoplay, Pagination, Navigation]}
								className="mySwiper"
							>
								{mainPageProducts.map((product) => {
									return (
										<SwiperSlide onClick={() => viewProductDetail(product)}>
											<img src={product.Image.Content} />
										</SwiperSlide>
									);
								})}
							</Swiper>
						</Col>
					</Row>
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
