// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"github.com/Duane-Arzu/test1/internal/data"
// 	_ "github.com/Duane-Arzu/test1/internal/data"
// 	"github.com/Duane-Arzu/test1/internal/validator"
// 	_ "github.com/Duane-Arzu/test1/internal/validator"
// )

<<<<<<< HEAD
// func (a *applicationDependences) createCommentHandler(w http.ResponseWriter, r *http.Request) {
// 	//create a struct to hold a comment
// 	//we use struct tags [` `] to make the names display in lowercase
// 	var incomingData struct {
// 		Content string `json:"content"`
// 		Author  string `json:"author"`
// 	}

// 	//perform decoding

// 	err := a.readJSON(w, r, &incomingData)
// 	if err != nil {
// 		a.badRequestResponse(w, r, err)
// 		return
// 	}

// 	comment := &data.Comment{
// 		Content: incomingData.Content,
// 		Author:  incomingData.Author,
// 	}

// 	v := validator.New()
// 	//do validation
// 	data.ValidateComment(v, comment)
// 	if !v.IsEmpty() {
// 		a.failedValidationResponse(w, r, v.Errors) //implemented later
// 		return
// 	}

// 	//add comment to the comments table in database
// 	err = a.commentModel.Insert(comment)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	//for now display the result
// 	// fmt.Fprintf(w, "%+v\n", incomingData)

// 	//set a location header, the path to the newly created comments
// 	headers := make(http.Header)
// 	headers.Set("Location", fmt.Sprintf("/v1/comments/%d", comment.ID))

// 	//send a json response with a 201 (new reseource created) status code
// 	data := envelope{
// 		"comment": comment,
// 	}
// 	err = a.writeJSON(w, http.StatusCreated, data, headers)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}
// }

// func (a *applicationDependences) fetchCommentByID(w http.ResponseWriter, r *http.Request) (*data.Comment, error) {
// 	// Get the id from the URL /v1/comments/:id so that we
// 	// can use it to query the comments table. We will
// 	// implement the readIDParam() function later
// 	id, err := a.readIDParam(r)
// 	if err != nil {
// 		a.notFoundResponse(w, r)

// 	}

// 	// Call Get() to retrieve the comment with the specified id
// 	comment, err := a.commentModel.Get(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			a.notFoundResponse(w, r)
// 		default:
// 			a.serverErrorResponse(w, r, err)
// 		}

// 	}
// 	return comment, nil
// }

// func (a *applicationDependences) displayCommentHandler(w http.ResponseWriter, r *http.Request) {

// 	comment, err := a.fetchCommentByID(w, r)
// 	if err != nil {
// 		return
// 	}
// 	// display the comment
// 	data := envelope{
// 		"comment": comment,
// 	}
// 	err = a.writeJSON(w, http.StatusOK, data, nil)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}
=======

var incomingData struct {
	Content *string `json:"content"`
	Author  *string `json:"author"`
}

func (a *applicationDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	// create a struct to hold a comment
	// we use struct tags to make the names display in lowercase
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}
	// perform the decoding
	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	comment := &data.Comment{
		Content: incomingData.Content,
		Author:  incomingData.Author,
	}
	// Initialize a Validator instance
	v := validator.New()

	data.ValidateComment(v, comment)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors) // implemented later
		return
	}
	err = a.commentModel.Insert(comment)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData) // delete this
	// Set a Location header. The path to the newly created comment
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/comments/%d", comment.ID))

	data := envelope{
		"comment": comment,
	}
	err = a.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// for now display the result
	fmt.Fprintf(w, "%+v\n", incomingData)
}

func (a *applicationDependencies) displayCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL /v1/comments/:id so that we
	// can use it to query teh comments table. We will
	// implement the readIDParam() function later
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	// Call Get() to retrieve the comment with the specified id
	comment, err := a.commentModel.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	// display the comment
	data := envelope{
		"comment": comment,
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
>>>>>>> 0cc1270f48216b9318fcc1ef24b827397488e322

// }

<<<<<<< HEAD
// func (a *applicationDependences) updateCommentHandler(w http.ResponseWriter, r *http.Request) {

// 	comment, err := a.fetchCommentByID(w, r)
// 	if err != nil {
// 		return
// 	}

// 	// Use our temporary incomingData struct to hold the data
// 	// Note: I have changed the types to pointer to differentiate
// 	// between the client leaving a field empty intentionally
// 	// and the field not needing to be updated
// 	var incomingData struct {
// 		Content *string `json:"content"`
// 		Author  *string `json:"author"`
// 	}

// 	// perform the decoding
// 	err = a.readJSON(w, r, &incomingData)
// 	if err != nil {
// 		a.badRequestResponse(w, r, err)
// 		return
// 	}
// 	// We need to now check the fields to see which ones need updating
// 	// if incomingData.Content is nil, no update was provided
// 	if incomingData.Content != nil {
// 		comment.Content = *incomingData.Content
// 	}
// 	// if incomingData.Author is nil, no update was provided
// 	if incomingData.Author != nil {
// 		comment.Author = *incomingData.Author
// 	}

// 	// Before we write the updates to the DB let's validate
// 	v := validator.New()
// 	data.ValidateComment(v, comment)
// 	if !v.IsEmpty() {
// 		a.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}

// 	// perform the update
// 	err = a.commentModel.Update(comment)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}
// 	data := envelope{
// 		"comment": comment,
// 	}
// 	err = a.writeJSON(w, http.StatusOK, data, nil)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}

// }

// func (a *applicationDependences) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := a.readIDParam(r)
// 	if err != nil {
// 		a.notFoundResponse(w, r)
// 		return
// 	}
// 	err = a.commentModel.Delete(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			a.notFoundResponse(w, r)
// 		default:
// 			a.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}

// 	//diplay the comment
// 	data := envelope{
// 		"message": "comment deleted successfully",
// 	}
// 	err = a.writeJSON(w, http.StatusOK, data, nil)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 	}
// }

// func (a *applicationDependences) listCommentHandler(w http.ResponseWriter, r *http.Request) {
// 	//create a struct to hold the query parameters
// 	//Later, fields will be added for pagination and sorting (filters)
// 	var queryParameterData struct {
// 		Content string
// 		Author  string
// 		data.Filters
// 	}

// 	//get query parameters from url
// 	queryParameter := r.URL.Query()

// 	//load the query parameters into the created struct
// 	queryParameterData.Content = a.getSingleQueryParameter(queryParameter, "content", "")
// 	queryParameterData.Author = a.getSingleQueryParameter(queryParameter, "author", "")

// 	// Create a new validator instance
// 	v := validator.New()

// 	queryParameterData.Filters.Page = a.getSingleIntegerParameter(queryParameter, "page", 1, v)
// 	queryParameterData.Filters.PageSize = a.getSingleIntegerParameter(queryParameter, "page_size", 10, v)

// 	// Check if our filters are valid
// 	data.ValidateFilters(v, queryParameterData.Filters)
// 	if !v.IsEmpty() {
// 		a.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}

// 	//call GetAll to retrieve all comments of the DB
// 	comments, err := a.commentModel.GetAll(queryParameterData.Content, queryParameterData.Author, queryParameterData.Filters)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			a.notFoundResponse(w, r)
// 			return
// 		default:
// 			a.serverErrorResponse(w, r, err)
// 			return
// 		}
// 	}

// 	data := envelope{
// 		"comments": comments,
// 	}
// 	err = a.writeJSON(w, http.StatusOK, data, nil)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}
// }
=======
func (a *applicationDependencies) updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	// Retrieve the comment from the database
	comment, err := a.commentModel.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	// Decode the incoming JSON
	err = a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	// Update the comment fields based on the incoming data
	if incomingData.Content != nil {
		comment.Content = *incomingData.Content
	}
	if incomingData.Author != nil {
		comment.Author = *incomingData.Author
	}

	// Validate the updated comment
	v := validator.New()
	data.ValidateComment(v, comment)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Perform the update in the database
	err = a.commentModel.Update(comment)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// Respond with the updated comment
	data := envelope{
		"comment": comment,
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *applicationDependencies) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	err = a.commentModel.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.IDnotFound(w, r, id) // Pass the ID to the custom message handler
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	data := envelope{
		"message": "comment successfully deleted",
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *applicationDependencies) listCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Create a struct to hold the query parameters
	// Later on we will add fields for pagination and sorting (filters)
	var queryParametersData struct {
		Content string
		Author  string
		data.Filters
	}
	// get the query parameters from the URL
	queryParameters := r.URL.Query()
	// Load the query parameters into our struct
	queryParametersData.Content = a.getSingleQueryParameter(
		queryParameters,
		"content",
		"")

	queryParametersData.Author = a.getSingleQueryParameter(
		queryParameters,
		"author",
		"")

	v := validator.New()
	queryParametersData.Filters.Page = a.getSingleIntegerParameter(
		queryParameters, "page", 1, v)
	queryParametersData.Filters.PageSize = a.getSingleIntegerParameter(
		queryParameters, "page_size", 10, v)

	queryParametersData.Filters.Sort = a.getSingleQueryParameter(
		queryParameters, "sort", "id")

	queryParametersData.Filters.SortSafeList = []string{"id", "author",
		"-id", "-author"}

	// Check if our filters are valid
	data.ValidateFilters(v, queryParametersData.Filters)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	comments, metadata, err := a.commentModel.GetAll(
		queryParametersData.Content,
		queryParametersData.Author,
		queryParametersData.Filters,
	)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	data := envelope{
		"comments":  comments,
		"@metadata": metadata,
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
>>>>>>> 0cc1270f48216b9318fcc1ef24b827397488e322
