import "../../Assets/css/purchase-history.css";
import { Button, Table, Jumbotron, Collapse } from "reactstrap";

const PurchaseDetail = (props) => {
	const onToggleItem = () => {
		props.toggleItem && props.toggleItem();
	};

	return (
		<>
			<Collapse isOpen={props.isOpen}>
				<Jumbotron className="purchase-detail-jumbotron">
					<Table className="purchase-detail-table">
						<tbody>
							<tr>
								<th>Name</th>
								<th>Price</th>
								<th>Quantity</th>
								<th>Amount</th>
							</tr>
							{props.purchase.ProductPurchaseDetail.map((purchaseDetail) => {
								return (
									<tr>
										<td>{purchaseDetail.ProductName}</td>
										<td>{purchaseDetail.ProductPrice} RSD</td>
										<td>{purchaseDetail.ProductQuantity}</td>
										<td>
											{Number(purchaseDetail.ProductPrice) *
												Number(purchaseDetail.ProductQuantity)}{" "}
											RSD
										</td>
									</tr>
								);
							})}
						</tbody>
						<tfooter>
							<Button
								style={{ marginTop: "15px" }}
								className="purchase-history-buttons"
								type="button"
								onClick={onToggleItem}
							>
								Hide items
							</Button>
						</tfooter>
					</Table>
				</Jumbotron>
			</Collapse>
		</>
	);
};

export default PurchaseDetail;
