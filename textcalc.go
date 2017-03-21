package go-ttd-textcalc

func text_to_number(s string) int {
    //TODO: conver s to lowercase
  switch s {
  case "zero":
      return 0
    case "one":
      return 1
    case "two":
      return 2
    case "three":
      return 3
    case "four":
      return 4
    case "five":
      return 5
    default:
      //TODO: use Errors correctly
      return -1
  }
}
