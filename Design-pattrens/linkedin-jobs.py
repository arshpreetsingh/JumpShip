import time
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# Set your LinkedIn login credentials and search keyword
email = "arsh840@gmail.com"
password = "Amber252556!@"
search_keyword = "software engineer"

# Initialize the browser driver (replace with the appropriate driver for your browser)
driver = webdriver.Firefox()
# Alternatively, you can use Edge (or any other browser) as follows:
# driver = webdriver.Edge()

# Login to LinkedIn
def login(email, password):
    driver.get("https://www.linkedin.com/login")
    time.sleep(2)

    email_input = driver.find_element(By.ID, "username")
    password_input = driver.find_element(By.ID, "password")

    email_input.send_keys(email)
    password_input.send_keys(password)
    password_input.submit()


def search_jobs(search_keyword):
    WebDriverWait(driver, 20).until(
        EC.visibility_of_element_located(
            (By.CSS_SELECTOR, "input.search-global-typeahead__input.always-show-placeholder"))
    )
    search_bar = driver.find_element(By.CSS_SELECTOR, "input.search-global-typeahead__input.always-show-placeholder")

    search_bar.click()  # Click on the search bar to ensure it's interactable
    search_bar.send_keys(search_keyword)
    search_bar.send_keys(Keys.ENTER)
    time.sleep(2)

    jobs_tab = driver.find_element(By.LINK_TEXT, "Jobs")
    jobs_tab.click()


# Apply to jobs
def apply_jobs():
    easy_apply_buttons = driver.find_elements(By.CSS_SELECTOR, "button[data-control-name='A_jobssearch_job_result_click_easyApply']")

    for button in easy_apply_buttons:
        button.click()
        time.sleep(1)

        submit_application = WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "button.jobs-easy-apply-form__submit-button"))
        )
        submit_application.click()

        # Wait for the application to be submitted and close the modal
        time.sleep(2)
        close_button = driver.find_element(By.CSS_SELECTOR, "button.artdeco-modal__dismiss")
        close_button.click()
        time.sleep(1)

# Main function
def main():
    login(email, password)
    search_jobs(search_keyword)
    apply_jobs()
    driver.quit()

if __name__ == "__main__":
    main()
