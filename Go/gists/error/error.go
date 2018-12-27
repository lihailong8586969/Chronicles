"import github.com/pkg/errors"

func canProceed(id int) (bool, error) {
  found, err := checkData(id)
  if err != nil {
    return false, errors.Wrap(err, "checkData failed”)
  }
  return found, nil
}