package fixture

func Label(item Item) string {
	switch item.Kind {
	case Digital:
		return "download"
	case Subscription:
		return "yearly plan"
	default:
		return "ships in 3-5 days"
	}
}
