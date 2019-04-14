# terraform v0.11 compatible

variable "app_id" {
  type = "string"
}

variable "api_key" {
  type = "string"
}

provider "algolia" {
	application_id = "${var.app_id}"
	api_key = "${var.api_key}"
}

resource "algolia_index" "main_index_1" {
	advanced_syntax = false
	allow_compression_of_integer_array = false
	allow_typos_on_numeric_tokens = false
	attribute_for_distinct = ""
	attributes_for_faceting = ["brand", "category"]
	attributes_to_highlight = []
	attributes_to_retrieve = []
	attributes_to_snippet = []
	custom_ranking = []
	disable_typo_tolerance_on_attributes = []
	disable_typo_tolerance_on_words = []
	highlight_post_tag = "</em>"
	highlight_pre_tag = "<em>"
	hits_per_page = 20
	max_facet_hits = 10
	max_values_per_facet = 100
	min_proximity = 1
	min_word_size_for_1_typo = 4
	min_word_size_for_2_typos = 8
	optional_words = []
	pagination_limited_to = 1000
	query_type = "prefixLast"
	ranking = []
	remove_words_if_no_results = "none"
	replace_synonyms_in_highlight = false
	response_fields = []
	restrict_highlight_and_snippet_arrays = false
	searchable_attributes = ["name", "brand_name", "category_name"]
	separators_to_index = ""
	snippet_ellipsis_text = ""
	typo_tolerance = "min"
	unretrievable_attributes = []
	name = "test-index.20190414"
}

# resource "algolia_index" "main_index_2" {
	# name = "test-index-2.20190414"
	# }
