import { Route } from "react-router-dom";
import * as videoGameAPI from "../../APIs/ProductMicroservice/video_game_api";
import * as consoleAPI from "../../APIs/ProductMicroservice/console_api";
import * as hddAPI from "../../APIs/ProductMicroservice/hard_disk_drive_api";
import * as headphonesAPI from "../../APIs/ProductMicroservice/headphones_api";
import * as keyboardAPI from "../../APIs/ProductMicroservice/keyboard_api";
import * as monitorAPI from "../../APIs/ProductMicroservice/monitor_api";
import * as motherboardAPI from "../../APIs/ProductMicroservice/motherboard_api";
import * as mouseAPI from "../../APIs/ProductMicroservice/mouse_api";
import * as processorAPI from "../../APIs/ProductMicroservice/processor_api";
import * as psuAPI from "../../APIs/ProductMicroservice/psu_api";
import * as ramAPI from "../../APIs/ProductMicroservice/ram_api";
import * as ssdAPI from "../../APIs/ProductMicroservice/solid_state_drive_api";
import * as graphicsCardsAPI from "../../APIs/ProductMicroservice/graphics_card_api";

import ProductDetail from "../../Pages/ProductDetail";

export const ProductDetailVideoGame = () => (
	<Route
		path="/videoGames/:id"
		element={
			<ProductDetail
				key="/videoGames/:id"
				GET_PRODUCT_BY_ID={videoGameAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailConsole = () => (
	<Route
		path="/consoles/:id"
		element={
			<ProductDetail
				key="/consoles/:id"
				GET_PRODUCT_BY_ID={consoleAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailGraphicsCard = () => (
	<Route
		path="/graphicsCards/:id"
		element={
			<ProductDetail
				key="/graphicsCards/:id"
				GET_PRODUCT_BY_ID={graphicsCardsAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailRAM = () => (
	<Route
		path="/rams/:id"
		element={
			<ProductDetail key="/rams/:id" GET_PRODUCT_BY_ID={ramAPI.GET_BY_ID} />
		}
	/>
);

export const ProductDetailMotherboard = () => (
	<Route
		path="/motherboards/:id"
		element={
			<ProductDetail
				key="/motherboards/:id"
				GET_PRODUCT_BY_ID={motherboardAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailProcessor = () => (
	<Route
		path="/processors/:id"
		element={
			<ProductDetail
				key="/processors/:id"
				GET_PRODUCT_BY_ID={processorAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailHDD = () => (
	<Route
		path="/hdds/:id"
		element={
			<ProductDetail key="/hdds/:id" GET_PRODUCT_BY_ID={hddAPI.GET_BY_ID} />
		}
	/>
);

export const ProductDetailSSD = () => (
	<Route
		key="/ssds/:id"
		path="/ssds/:id"
		element={
			<ProductDetail key="/ssds/:id" GET_PRODUCT_BY_ID={ssdAPI.GET_BY_ID} />
		}
	/>
);

export const ProductDetailMonitor = () => (
	<Route
		path="/monitors/:id"
		element={
			<ProductDetail
				key="/monitors/:id"
				GET_PRODUCT_BY_ID={monitorAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailPSU = () => (
	<Route
		path="/psus/:id"
		element={
			<ProductDetail key="/psus/:id" GET_PRODUCT_BY_ID={psuAPI.GET_BY_ID} />
		}
	/>
);

export const ProductDetailHeadphones = () => (
	<Route
		path="/headphones/:id"
		element={
			<ProductDetail
				key="/headphones/:id"
				GET_PRODUCT_BY_ID={headphonesAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailKeyboard = () => (
	<Route
		path="/keyboards/:id"
		element={
			<ProductDetail
				key="/keyboards/:id"
				GET_PRODUCT_BY_ID={keyboardAPI.GET_BY_ID}
			/>
		}
	/>
);

export const ProductDetailMouse = () => (
	<Route
		path="/mice/:id"
		element={
			<ProductDetail key="/mice/:id" GET_PRODUCT_BY_ID={mouseAPI.GET_BY_ID} />
		}
	/>
);
