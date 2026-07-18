# GOV.UK Design System Coverage

This repo embeds GOV.UK Frontend assets and provides Go template wrappers for GOV.UK Design System components and patterns.

## Styles

Styles are provided by the embedded GOV.UK Frontend distribution (CSS, JS and images) under `internal/govuk/assets/build/govuk-frontend`.

## Components

The following GOV.UK Design System components have Go template implementations under `internal/govuk/templates/components` and public Go wrappers under `pkg/govuk/components`.

- accordion
- back-link (`back_link`)
- breadcrumbs
- button
- character-count (`character_count`)
- checkboxes
- cookie-banner (`cookie_banner`)
- date-input (`date_input`)
- details
- error-message (`error_message`)
- error-summary (`error_summary`)
- exit-this-page (`exit_this_page`)
- fieldset
- generic-header (`generic_header`)
- file-upload (`file_upload`)
- footer
- header
- hint
- inset-text (`inset_text`)
- label
- notification-banner (`notification_banner`)
- pagination
- panel
- password-input (`password_input`)
- phase-banner (`phase_banner`)
- radios
- select (`select_`)
- service-navigation (`service_navigation`)
- skip-link (`skip_link`)
- summary-list (`summary_list`)
- table
- tabs
- tag
- task-list (`task_list`)
- text-input (`text_input`)
- textarea
- warning-text (`warning_text`)

## Patterns

The GOV.UK Design System patterns are available as templates under `internal/govuk/templates/patterns` and public wrappers under `pkg/govuk/patterns`.

- addresses
- bank-details (`bank_details`)
- check-a-service-is-suitable (`check_a_service_is_suitable`)
- check-answers (`check_answers`)
- complete-multiple-tasks (`complete_multiple_tasks`)
- confirm-a-phone-number (`confirm_a_phone_number`)
- confirm-an-email-address (`confirm_an_email_address`)
- confirmation-pages (`confirmation_pages`)
- contact-a-department-or-service-team (`contact_a_department_or_service_team`)
- cookies-page (`cookies_page`)
- create-a-username (`create_a_username`)
- create-accounts (`create_accounts`)
- dates
- email-addresses (`email_addresses`)
- equality-information (`equality_information`)
- exit-a-page-quickly (`exit_a_page_quickly`)
- names
- national-insurance-numbers (`national_insurance_numbers`)
- navigate-a-service (`navigate_a_service`)
- page-not-found-pages (`page_not_found_pages`)
- passwords
- payment-card-details (`payment_card_details`)
- phone-numbers (`phone_numbers`)
- question-pages (`question_pages`)
- recover-from-validation-errors (`recover_from_validation_errors`)
- service-unavailable-pages (`service_unavailable_pages`)
- start-using-a-service (`start_using_a_service`)
- step-by-step-navigation (`step_by_step_navigation`)
- there-is-a-problem-with-the-service-pages (`there_is_a_problem_with_the_service_pages`)
